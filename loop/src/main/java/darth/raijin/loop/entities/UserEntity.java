package darth.raijin.loop.entities;

import java.util.Set;
import java.util.UUID;

import jakarta.persistence.*;
import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.NotBlank;

@Entity
@Table(name = "users")
public class UserEntity {

  @Id
  @GeneratedValue(strategy = GenerationType.UUID)
  private UUID id;

  @Column(name = "name")
  @NotBlank
  private String name;

  @Column(name = "username", unique = true)
  private String username;

  @Column(name = "email", unique = true)
  @NotBlank
  @Email
  private String email;

  @Column(name = "password")
  private String password;

  @ManyToMany(fetch = FetchType.LAZY, cascade = CascadeType.ALL)
  @JoinTable(
          name = "user_hosted_events",
          joinColumns = @JoinColumn(name = "host_id"),
          inverseJoinColumns = @JoinColumn(name = "event_id")
  )
  private Set<EventEntity> hostedEvents;

  @ManyToMany(fetch = FetchType.LAZY, cascade = CascadeType.ALL)
  @JoinTable(
          name = "user_participated_events",
          joinColumns = @JoinColumn(name = "attendee_id"),
          inverseJoinColumns = @JoinColumn(name = "event_id")
  )
  private Set<EventEntity> participatedEvents;

  /** Constructor used for registering user */
  public UserEntity(String name, String username, String email, String password) {
    this.name = name;
    this.username = username;
    this.email = email;
    this.password = password;
  }

    public UserEntity() {

    }

    public UUID getId() {
    return id;
  }

  public void setId(UUID id) {
    this.id = id;
  }

  public String getName() {
    return name;
  }

  public void setName(String name) {
    this.name = name;
  }

  public String getUsername() {
    return username;
  }

  public void setUsername(String username) {
    this.username = username;
  }

  public String getEmail() {
    return email;
  }

  public void setEmail(String email) {
    this.email = email;
  }

  public String getPassword() {
    return password;
  }

  public void setPassword(String password) {
    this.password = password;
  }
}
